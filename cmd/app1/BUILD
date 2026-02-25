go_package()

go_binary(
    name="bin",
)

docker_image(
    name="docker",
    repository="ghcr.io/treilly94/pants/app1",
    dependencies=[":bin"],
    image_tags=["latest"]
)
