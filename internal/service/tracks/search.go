package tracks

import (
	"context"
	"github.com/mdafaardiansyah/musicacu-backend/internal/models/spotify"
	"github.com/mdafaardiansyah/musicacu-backend/internal/models/trackactivities"
	spotifyRepo "github.com/mdafaardiansyah/musicacu-backend/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

// Search will search the track based on given query and will return the track that match
// with the query. The search result will be paginated and the pagination size can be
// controlled by pageSize parameter. The start index of pagination can be controlled by
// pageIndex parameter.
//
// The method will also return the user's track activities based on the trackID that
// returned in search result.
func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.spotifyOutbound.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error search track to spotify")
		return nil, err
	}
	trackIDs := make([]string, len(trackDetails.Tracks.Items))
	for idx, item := range trackDetails.Tracks.Items {
		trackIDs[idx] = item.ID
	}

	trackActivities, err := s.trackActivitiesRepo.GetBulkSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("error get track activities from database")
		return nil, err
	}

	return modelToResponse(trackDetails, trackActivities), nil
}

// modelToResponse transforms a SpotifySearchResponse and a map of track activities
// into a SearchResponse. It extracts essential track, album, and artist details,
// while also incorporating user-specific track activity data, such as like status.
// If the input data is nil, it returns nil.
func modelToResponse(data *spotifyRepo.SpotifySearchResponse, mapTrackActivities map[string]trackactivities.TrackActivity) *spotify.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks.Items {
		artistsName := make([]string, len(item.Artists))
		for idx, artist := range item.Artists {
			artistsName[idx] = artist.Name
		}

		imageUrls := make([]string, len(item.Album.Images))
		for idx, image := range item.Album.Images {
			imageUrls[idx] = image.URL
		}

		items = append(items, spotify.SpotifyTrackObject{
			// album related fields
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesURL:   imageUrls,
			AlbumName:        item.Album.Name,
			// artist related fields
			ArtistsName: artistsName,
			// track related fields
			Explicit: item.Explicit,
			ID:       item.ID,
			Name:     item.Name,
			// track activity isliked
			IsLiked: mapTrackActivities[item.ID].IsLiked,
		})
	}

	return &spotify.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Items:  items,
		Total:  data.Tracks.Total,
	}
}
