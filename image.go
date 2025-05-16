package mcsmapi

import (
	"log"

	"github.com/google/go-querystring/query"
)

// GetImageListRequest defines the query parameters for getting a list of images.
type GetImageListRequest struct {
	DaemonID string `url:"daemonId"`
}

func (q *GetImageListRequest) BuildQueryString() string {
	v, err := query.Values(q)
	if err != nil {
		log.Fatal(err)
	}
	return v.Encode()
}

// GetImageListResponse defines the response structure for getting a list of images.
type GetImageListResponse struct {
	Data DockerImageList `json:"data"`
	Time int64           `json:"time"`
}

// DockerImageList defines a list of Docker images.
type DockerImageList struct {
	Images []DockerImage `json:"images"` // List of Docker images
}

// DockerImage defines an individual Docker image.
type DockerImage struct {
	RepoTags []string `json:"repoTags"` // List of repository tags for the image
	Size     int64    `json:"size"`     // Image size in bytes
	Created  string   `json:"created"`  // Creation date of the image
	ID       string   `json:"id"`       // Image ID
}

type GetContainerListResponse struct {
	Data DockerContainerList `json:"data"`
	Time int64               `json:"time"`
}

// DockerContainerList defines a list of Docker containers.
type DockerContainerList struct {
	Containers []DockerContainer `json:"containers"`
}

// DockerContainer defines an individual Docker container.
type DockerContainer struct {
	ID      string   `json:"id"`
	Names   []string `json:"names"`
	Image   string   `json:"image"`
	ImageID string   `json:"imageID"`
	State   string   `json:"state"`
	Status  string   `json:"status"`
}

// GetNetworkListResponse defines the response structure for getting a list of networks.
type GetNetworkListResponse struct {
	Data DockerNetworkList `json:"data"`
	Time int64             `json:"time"`
}

// DockerNetworkList defines a list of Docker networks.
type DockerNetworkList struct {
	Networks []DockerNetwork `json:"networks"`
}

// DockerNetwork defines an individual Docker network.
type DockerNetwork struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Driver string   `json:"driver"`
	Scope  string   `json:"scope"`
	IPAM   any      `json:"ipam"`
	Labels any      `json:"labels"`
	Peers  []string `json:"peers"`
}

// CreateImageRequest defines the request body for creating a Docker image.
type CreateImageRequest struct {
	DockerFile string `json:"dockerFile"`
	Name       string `json:"name"`
	Tag        string `json:"tag"`
}

// GetBuildProgressResponse defines the response structure for build progress.
type GetBuildProgressResponse struct {
	Status int            `json:"status"`
	Data   map[string]int `json:"data"` // image: status (-1 = Failed, 1 = Building, 2 = Complete)
	Time   int64          `json:"time"`
}
