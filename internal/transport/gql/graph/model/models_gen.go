// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateNft struct {
	Name string `json:"name"`
}

type GetNft struct {
	ID string `json:"id"`
}

type Nft struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type UpdateNft struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
