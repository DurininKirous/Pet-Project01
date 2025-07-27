package usersHandlers


import (
    "encoding/json"
    "net/http"
	"strconv"
    "github.com/go-chi/chi/v5"
    users "project01/app/internal/service/users"
)

type Handler struct {
    service *users.Service
}

func New(service *users.Service) *Handler {
    return &Handler{service: service}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    users, err := h.service.GetAll(ctx)
    if err != nil {
        http.Error(w, "failed to get users", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}


func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

    user, err := h.service.GetById(ctx, id)
    if err != nil {
        http.Error(w, "user not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
 
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    var input users.CreateUserInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "invalid JSON body", http.StatusBadRequest)
        return
    }

    user, err := h.service.Create(ctx, input)
    if err != nil {
        http.Error(w, "failed to create user", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

