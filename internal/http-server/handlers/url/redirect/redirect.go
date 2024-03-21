package redirect

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage"
)

type URLGetter interface {
	GetURL(alias string) (string, error)
}

func New(logs *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.redirect.New"

		log := logs.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, resp.Error("not found"))

			return
		}

		resURL, err := urlGetter.GetURL(alias)
		if errors.Is(err, storage.ErrURLNotFOund) {
			log.Info("url not found", "alias", alias)

			render.JSON(w, r, resp.Error("not found"))
			return
		}
		if err != nil {
			log.Error("falied to get url", sl.Err(err))
			render.JSON(w, r, resp.Error("internal error"))
			return
		}

		log.Info("git url", slog.String("url", resURL))

		http.Redirect(w, r, resURL, http.StatusFound)
	}
}
