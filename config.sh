AUTH=$(echo -n "zhukovskiy:T3zrotWqKT" | base64)
cat << EOF > config.json
{
    "auths": {
        "https://index.docker.io/v1/": {
            "auth": "${AUTH}"
        }
    }
}
EOF