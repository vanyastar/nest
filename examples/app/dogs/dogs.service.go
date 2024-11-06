// An example of how to access from dogs to cats service

package dogs

import "example/app/cats"

type DogsService struct {
	catsService *cats.CatsService
}

var GetDogsService = &DogsService{
	catsService: cats.GetCatsService,
}

func (this *DogsService) GetCatName() string {
	return this.catsService.GetName()
}
