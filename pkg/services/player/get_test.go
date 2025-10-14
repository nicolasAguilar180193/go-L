package player_test

import (
	"errors"
	"testing"

	"github.com/nicolasAguilar180193/go-L/mocks"
	"github.com/nicolasAguilar180193/go-L/pkg/domain"
	"github.com/nicolasAguilar180193/go-L/pkg/services/player"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Get(t *testing.T) {

	// Define test cases
	testTable := map[string]struct {
		playerId   string                                                     // input
		setup      func(mockRepo *mocks.MockPlayerRepository)                 // setup function to configure mock behavior
		assertFunc func(subTest *testing.T, player *domain.Player, err error) // assertion function
	}{
		"generic error": {
			playerId: "any-id",
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Get("any-id").Return(nil, errors.New("generic error"))
			},
			assertFunc: func(subTest *testing.T, player *domain.Player, err error) {
				assert.Nil(subTest, player)
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error getting player")
			},
		},
		"successful retrieval": {
			playerId: "valid-id",
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Get("valid-id").Return(&domain.Player{
					ID: "valid-id",
				}, nil)
			},
			assertFunc: func(subTest *testing.T, player *domain.Player, err error) {
				assert.NotNil(subTest, player)
				assert.Equal(subTest, "valid-id", player.ID)
				assert.NoError(subTest, err)
			},
		},
		"empty id": {
			playerId: "",
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				// No setup needed as the function should return early
			},
			assertFunc: func(subTest *testing.T, player *domain.Player, err error) {
				assert.Nil(subTest, player)
				assert.NotNil(subTest, err)
				assert.Equal(subTest, "player ID is required", err.Error())
			},
		},
		"not found error": {
			playerId: "missing-id",
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Get("missing-id").Return(nil, domain.ErrNotFound)
			},
			assertFunc: func(subTest *testing.T, player *domain.Player, err error) {
				assert.Nil(subTest, player)
				assert.NotNil(subTest, err)

				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeNotFound, appErr.Code)
					assert.Equal(subTest, "player with id 'missing-id' not found", appErr.Msg)
				} else {
					t.Errorf("expected error of type AppError, got %T", err)
				}
			},
		},
		"timeout": {
			playerId: "timeout-id",
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Get("timeout-id").Return(nil, domain.ErrTimeout)
			},
			assertFunc: func(subTest *testing.T, player *domain.Player, err error) {
				assert.Nil(subTest, player)
				assert.NotNil(subTest, err)

				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeTimeout, appErr.Code)
					assert.Equal(subTest, "timeout error, try again later", appErr.Msg)
				} else {
					t.Errorf("expected error of type AppError, got %T", err)
				}
			},
		},
	}

	// Run tests
	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {
			ctrl := gomock.NewController(subTest)
			mockPlayerRepo := mocks.NewMockPlayerRepository(ctrl)

			s := &player.Service{
				Repository: mockPlayerRepo,
			}

			test.setup(mockPlayerRepo)

			player, err := s.Get(test.playerId)

			test.assertFunc(subTest, player, err)
		})
	}
}
