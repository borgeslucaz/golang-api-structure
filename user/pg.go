package user

import(
	"github.com/go-pg/pg"
	"github.com/borgeslucaz/golang-api-structure/models"
	
)
type repo struct {
	db *pg.DB
}

// StartRepository postgers respo
func StartRepository(p *pg.DB) Repository {
	return &repo{
		db: p,
	}
}

func (r *repo) Find(id int) (*models.User, error) {
	user := &models.User{Base:models.Base{ID:id}}
    err := r.db.Select(user)
    if err != nil {
        return user, err
	}
	return user, err
}

func (r *repo) FindByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := r.db.Model(user).Column("*").Where("email = ?", email).Select()
    if err != nil {
        return nil, err
	}
	return user, err
}

// Create persist the user on database
func (r *repo) Create(user *models.User) (*models.User, error) {
	err := r.db.Insert(user)
    if err != nil {
        panic(err)
	}
	return user, err
}

// func (r *repo) FindAll() ([]*User, error) {
// }

// func (r *repo) Update(user *User) error {
// }

// func (r *repo) Store(user *User) (entity.ID, error) {
// }