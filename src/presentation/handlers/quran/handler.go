package quran

import (
	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/query"
)

// Handler contains all handlers for Quran operations
type Handler struct {
	// Surah handlers
	createSurahHandler      *CreateSurahHandler
	updateSurahHandler      *UpdateSurahHandler
	deleteSurahHandler      *DeleteSurahHandler
	getSurahByIdHandler     *GetSurahByIdHandler
	getSurahByNumberHandler *GetSurahByNumberHandler
	getAllSurahsHandler     *GetAllSurahsHandler

	// Ayah handlers
	createAyahHandler        *CreateAyahHandler
	updateAyahHandler        *UpdateAyahHandler
	deleteAyahHandler        *DeleteAyahHandler
	getAyahByIdHandler       *GetAyahByIdHandler
	getAyahsBySurahHandler   *GetAyahsBySurahHandler
	getAyahsByJuzHandler     *GetAyahsByJuzHandler

	// Juz handlers
	createJuzHandler      *CreateJuzHandler
	updateJuzHandler      *UpdateJuzHandler
	deleteJuzHandler      *DeleteJuzHandler
	getJuzByIdHandler     *GetJuzByIdHandler
	getJuzByNumberHandler *GetJuzByNumberHandler
	getAllJuzHandler      *GetAllJuzHandler
}

// NewHandler creates a new Quran handler
func NewHandler(
	commandHandler *command.CommandHandler,
	queryHandler *query.QueryHandler,
) *Handler {
	return &Handler{
		// Surah handlers
		createSurahHandler:      NewCreateSurahHandler(commandHandler),
		updateSurahHandler:      NewUpdateSurahHandler(commandHandler),
		deleteSurahHandler:      NewDeleteSurahHandler(commandHandler),
		getSurahByIdHandler:     NewGetSurahByIdHandler(queryHandler),
		getSurahByNumberHandler: NewGetSurahByNumberHandler(queryHandler),
		getAllSurahsHandler:     NewGetAllSurahsHandler(queryHandler),

		// Ayah handlers
		createAyahHandler:        NewCreateAyahHandler(commandHandler),
		updateAyahHandler:        NewUpdateAyahHandler(commandHandler),
		deleteAyahHandler:        NewDeleteAyahHandler(commandHandler),
		getAyahByIdHandler:       NewGetAyahByIdHandler(queryHandler),
		getAyahsBySurahHandler:   NewGetAyahsBySurahHandler(queryHandler),
		getAyahsByJuzHandler:     NewGetAyahsByJuzHandler(queryHandler),

		// Juz handlers
		createJuzHandler:      NewCreateJuzHandler(commandHandler),
		updateJuzHandler:      NewUpdateJuzHandler(commandHandler),
		deleteJuzHandler:      NewDeleteJuzHandler(commandHandler),
		getJuzByIdHandler:     NewGetJuzByIdHandler(queryHandler),
		getJuzByNumberHandler: NewGetJuzByNumberHandler(queryHandler),
		getAllJuzHandler:      NewGetAllJuzHandler(queryHandler),
	}
}

// Surah handler getters
func (h *Handler) CreateSurahHandler() *CreateSurahHandler {
	return h.createSurahHandler
}

func (h *Handler) UpdateSurahHandler() *UpdateSurahHandler {
	return h.updateSurahHandler
}

func (h *Handler) DeleteSurahHandler() *DeleteSurahHandler {
	return h.deleteSurahHandler
}

func (h *Handler) GetSurahByIdHandler() *GetSurahByIdHandler {
	return h.getSurahByIdHandler
}

func (h *Handler) GetSurahByNumberHandler() *GetSurahByNumberHandler {
	return h.getSurahByNumberHandler
}

func (h *Handler) GetAllSurahsHandler() *GetAllSurahsHandler {
	return h.getAllSurahsHandler
}

// Ayah handler getters
func (h *Handler) CreateAyahHandler() *CreateAyahHandler {
	return h.createAyahHandler
}

func (h *Handler) UpdateAyahHandler() *UpdateAyahHandler {
	return h.updateAyahHandler
}

func (h *Handler) DeleteAyahHandler() *DeleteAyahHandler {
	return h.deleteAyahHandler
}

func (h *Handler) GetAyahByIdHandler() *GetAyahByIdHandler {
	return h.getAyahByIdHandler
}

func (h *Handler) GetAyahsBySurahHandler() *GetAyahsBySurahHandler {
	return h.getAyahsBySurahHandler
}

func (h *Handler) GetAyahsByJuzHandler() *GetAyahsByJuzHandler {
	return h.getAyahsByJuzHandler
}

// Juz handler getters
func (h *Handler) CreateJuzHandler() *CreateJuzHandler {
	return h.createJuzHandler
}

func (h *Handler) UpdateJuzHandler() *UpdateJuzHandler {
	return h.updateJuzHandler
}

func (h *Handler) DeleteJuzHandler() *DeleteJuzHandler {
	return h.deleteJuzHandler
}

func (h *Handler) GetJuzByIdHandler() *GetJuzByIdHandler {
	return h.getJuzByIdHandler
}

func (h *Handler) GetJuzByNumberHandler() *GetJuzByNumberHandler {
	return h.getJuzByNumberHandler
}

func (h *Handler) GetAllJuzHandler() *GetAllJuzHandler {
	return h.getAllJuzHandler
}