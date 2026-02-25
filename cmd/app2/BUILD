go_package()

go_binary(
    name="bin",
)

docker_image(
    name="docker",
    repository="ghcr.io/treilly94/pants/app2",
    image_tags=["latest"],
    dependencies=[":bin", "cmd/app1:docker"],
)
