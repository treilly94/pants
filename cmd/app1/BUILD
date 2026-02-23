go_package()

go_binary(
    name="bin",
)

docker_image(
    name="docker",
    repository="app1",
    dependencies=[":bin"],
    image_tags=["dev", "latest"]
)
