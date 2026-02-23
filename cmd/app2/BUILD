go_package()

go_binary(
    name="bin",
)

docker_image(
    name="docker",
    repository="app2",
    image_tags=["dev", "latest"],
    dependencies=[":bin", "cmd/app1:docker"],
)
