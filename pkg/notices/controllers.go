package notices

import (
	"BrickPlus/config"
	"BrickPlus/database/dbmodel"
	"BrickPlus/pkg/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"strconv"
)

type NoticesConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *NoticesConfigurator {
	return &NoticesConfigurator{configuration}
}

func noticeToModel(notices []*dbmodel.Notice) []models.Notice {
	noticesEdited := []models.Notice{}
	for _, notice := range notices {
		noticeEdited := models.Notice{
			IdUser:    notice.IdUser,
			IdProduct: notice.IdProduct,
			Content:   notice.Content,
		}
		noticesEdited = append(noticesEdited, noticeEdited)
	}
	return noticesEdited
}

func (config *NoticesConfigurator) noticeByIdHandler(w http.ResponseWriter, r *http.Request) {
	noticeId := chi.URLParam(r, "id")
	notices, err := config.NoticesRepository.FindNoticeById(noticeId)
	if err != nil || len(notices) == 0 {
		render.JSON(w, r, map[string]string{"error": "Notice not found"})
		return
	}
	render.JSON(w, r, noticeToModel(notices))
}

func (config *NoticesConfigurator) noticesByProductHandler(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")
	notices, err := config.NoticesRepository.FindNoticesByProductId(productId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to load notices for this product"})
		return
	}
	render.JSON(w, r, noticeToModel(notices))
}

func (config *NoticesConfigurator) noticesByUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	notices, err := config.NoticesRepository.FindNoticesByUserId(userId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to load notices for this user"})
		return
	}
	render.JSON(w, r, noticeToModel(notices))
}

func (config *NoticesConfigurator) addNoticeHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Notice{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	// Existence checks
	user, err := config.UsersRepository.FindUserById(strconv.Itoa(req.IdUser))
	if err != nil || len(user) == 0 {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}
	product, err := config.ProductsRepository.FindProductById(strconv.Itoa(req.IdProduct))
	if err != nil || len(product) == 0 {
		render.JSON(w, r, map[string]string{"error": "Product not found"})
		return
	}

	newNotice := &dbmodel.Notice{
		IdUser:    req.IdUser,
		IdProduct: req.IdProduct,
		Content:   req.Content,
	}
	_, err = config.NoticesRepository.Create(newNotice)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create notice"})
		return
	}
	render.JSON(w, r, map[string]string{"success": "Notice successfully created"})
}

func (config *NoticesConfigurator) deleteNoticeHandler(w http.ResponseWriter, r *http.Request) {
	noticeId := chi.URLParam(r, "id")
	notices, err := config.NoticesRepository.FindNoticeById(noticeId)
	if err != nil || len(notices) == 0 {
		render.JSON(w, r, map[string]string{"error": "Notice not found"})
		return
	}
	if err := config.NoticesRepository.Delete(notices[0]); err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete notice"})
		return
	}
	render.JSON(w, r, map[string]string{"success": "Notice successfully deleted"})
}

func (config *NoticesConfigurator) editNoticeHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Notice{}
	noticeId := chi.URLParam(r, "id")
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	// Existence check
	notices, err := config.NoticesRepository.FindNoticeById(noticeId)
	if err != nil || len(notices) == 0 {
		render.JSON(w, r, map[string]string{"error": "Notice not found"})
		return
	}

	updatedNotice := &dbmodel.Notice{
		Content: req.Content,
	}
	err = config.NoticesRepository.Update(updatedNotice, noticeId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update notice"})
		return
	}
	render.JSON(w, r, map[string]string{"success": "Notice successfully updated"})
}