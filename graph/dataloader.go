package graph

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ahmedkhalaf1996/GRQGoTaskx/models"

	"gorm.io/gorm"
)

// ------------- user loader start-------------------- //
const userloaderKey = "userloader"

func UserDataloaderMiddleware(DB *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userloader := UserLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []string) ([]*models.User, []error) {
					var users []*models.User
					err := DB.Where("id in (?)", ids).Find(&users).Error

					if err != nil {
						return nil, []error{err}
					}

					u := make(map[string]*models.User, len(users))

					for _, user := range users {
						u[fmt.Sprint(user.ID)] = user
					}

					result := make([]*models.User, len(ids))

					for i, id := range ids {
						result[i] = u[id]
					}

					return result, nil
				},
			}

			ctx := context.WithValue(r.Context(), userloaderKey, &userloader)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userloaderKey).(*UserLoader)
}

// ------------- user loader end-------------------- //
