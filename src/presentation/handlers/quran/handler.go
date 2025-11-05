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
	createAyahHandler      *CreateAyahHandler
	updateAyahHandler      *UpdateAyahHandler
	deleteAyahHandler      *DeleteAyahHandler
	getAyahByIdHandler     *GetAyahByIdHandler
	getAyahsBySurahHandler *GetAyahsBySurahHandler
	getAyahsByJuzHandler   *GetAyahsByJuzHandler

	// Juz handlers
	createJuzHandler      *CreateJuzHandler
	updateJuzHandler      *UpdateJuzHandler
	deleteJuzHandler      *DeleteJuzHandler
	getJuzByIdHandler     *GetJuzByIdHandler
	getJuzByNumberHandler *GetJuzByNumberHandler
	getAllJuzHandler      *GetAllJuzHandler

	// Translation handlers
	getTranslationEditionsHandler               *GetTranslationEditionsHandler
	getAyahTranslationByAyahAndEditionHandler   *GetAyahTranslationByAyahAndEditionHandler
	getAyahTranslationsBySurahAndEditionHandler *GetAyahTranslationsBySurahAndEditionHandler

	// Reciter handlers
	createReciterHandler    *CreateReciterHandler
	updateReciterHandler    *UpdateReciterHandler
	deleteReciterHandler    *DeleteReciterHandler
	getReciterByIdHandler   *GetReciterByIdHandler
	getReciterByNameHandler *GetReciterByNameHandler
	getAllRecitersHandler   *GetAllRecitersHandler

	// Ayah Audio File handlers
	createAyahAudioFileHandler                *CreateAyahAudioFileHandler
	updateAyahAudioFileHandler                *UpdateAyahAudioFileHandler
	deleteAyahAudioFileHandler                *DeleteAyahAudioFileHandler
	getAyahAudioFileByIdHandler               *GetAyahAudioFileByIdHandler
	getAyahAudioFileByAyahAndReciterHandler   *GetAyahAudioFileByAyahAndReciterHandler
	getAyahAudioFilesBySurahAndReciterHandler *GetAyahAudioFilesBySurahAndReciterHandler
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
		createAyahHandler:      NewCreateAyahHandler(commandHandler),
		updateAyahHandler:      NewUpdateAyahHandler(commandHandler),
		deleteAyahHandler:      NewDeleteAyahHandler(commandHandler),
		getAyahByIdHandler:     NewGetAyahByIdHandler(queryHandler),
		getAyahsBySurahHandler: NewGetAyahsBySurahHandler(queryHandler),
		getAyahsByJuzHandler:   NewGetAyahsByJuzHandler(queryHandler),

		// Juz handlers
		createJuzHandler:      NewCreateJuzHandler(commandHandler),
		updateJuzHandler:      NewUpdateJuzHandler(commandHandler),
		deleteJuzHandler:      NewDeleteJuzHandler(commandHandler),
		getJuzByIdHandler:     NewGetJuzByIdHandler(queryHandler),
		getJuzByNumberHandler: NewGetJuzByNumberHandler(queryHandler),
		getAllJuzHandler:      NewGetAllJuzHandler(queryHandler),

		// Translation handlers
		getTranslationEditionsHandler:               NewGetTranslationEditionsHandler(queryHandler),
		getAyahTranslationByAyahAndEditionHandler:   NewGetAyahTranslationByAyahAndEditionHandler(queryHandler),
		getAyahTranslationsBySurahAndEditionHandler: NewGetAyahTranslationsBySurahAndEditionHandler(queryHandler),

		// Reciter handlers
		createReciterHandler:    NewCreateReciterHandler(commandHandler),
		updateReciterHandler:    NewUpdateReciterHandler(commandHandler),
		deleteReciterHandler:    NewDeleteReciterHandler(commandHandler),
		getReciterByIdHandler:   NewGetReciterByIdHandler(queryHandler),
		getReciterByNameHandler: NewGetReciterByNameHandler(queryHandler),
		getAllRecitersHandler:   NewGetAllRecitersHandler(queryHandler),

		// Ayah Audio File handlers
		createAyahAudioFileHandler:                NewCreateAyahAudioFileHandler(commandHandler),
		updateAyahAudioFileHandler:                NewUpdateAyahAudioFileHandler(commandHandler),
		deleteAyahAudioFileHandler:                NewDeleteAyahAudioFileHandler(commandHandler),
		getAyahAudioFileByIdHandler:               NewGetAyahAudioFileByIdHandler(queryHandler),
		getAyahAudioFileByAyahAndReciterHandler:   NewGetAyahAudioFileByAyahAndReciterHandler(queryHandler),
		getAyahAudioFilesBySurahAndReciterHandler: NewGetAyahAudioFilesBySurahAndReciterHandler(queryHandler),
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

// Translation handler getters
func (h *Handler) GetTranslationEditionsHandler() *GetTranslationEditionsHandler {
	return h.getTranslationEditionsHandler
}

func (h *Handler) GetAyahTranslationByAyahAndEditionHandler() *GetAyahTranslationByAyahAndEditionHandler {
	return h.getAyahTranslationByAyahAndEditionHandler
}

func (h *Handler) GetAyahTranslationsBySurahAndEditionHandler() *GetAyahTranslationsBySurahAndEditionHandler {
	return h.getAyahTranslationsBySurahAndEditionHandler
}

// Reciter handler getters
func (h *Handler) CreateReciterHandler() *CreateReciterHandler   { return h.createReciterHandler }
func (h *Handler) UpdateReciterHandler() *UpdateReciterHandler   { return h.updateReciterHandler }
func (h *Handler) DeleteReciterHandler() *DeleteReciterHandler   { return h.deleteReciterHandler }
func (h *Handler) GetReciterByIdHandler() *GetReciterByIdHandler { return h.getReciterByIdHandler }
func (h *Handler) GetReciterByNameHandler() *GetReciterByNameHandler {
	return h.getReciterByNameHandler
}
func (h *Handler) GetAllRecitersHandler() *GetAllRecitersHandler { return h.getAllRecitersHandler }

// Ayah Audio File handler getters
func (h *Handler) CreateAyahAudioFileHandler() *CreateAyahAudioFileHandler {
	return h.createAyahAudioFileHandler
}
func (h *Handler) UpdateAyahAudioFileHandler() *UpdateAyahAudioFileHandler {
	return h.updateAyahAudioFileHandler
}
func (h *Handler) DeleteAyahAudioFileHandler() *DeleteAyahAudioFileHandler {
	return h.deleteAyahAudioFileHandler
}
func (h *Handler) GetAyahAudioFileByIdHandler() *GetAyahAudioFileByIdHandler {
	return h.getAyahAudioFileByIdHandler
}
func (h *Handler) GetAyahAudioFileByAyahAndReciterHandler() *GetAyahAudioFileByAyahAndReciterHandler {
	return h.getAyahAudioFileByAyahAndReciterHandler
}
func (h *Handler) GetAyahAudioFilesBySurahAndReciterHandler() *GetAyahAudioFilesBySurahAndReciterHandler {
	return h.getAyahAudioFilesBySurahAndReciterHandler
}
