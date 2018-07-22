# personal-playlist

This started as an idea for a personal offline youtube playlist possibly linking
with VLC to live update the playlist. The use case was never particularly
compelling, and fell apart. This is a quick trial RPC server that pretends to be
a small part of that project.

Neat commands:

    go build
    ./personal-playlist server
    ./personal-playlist add "asdf"
    ./personal-playlist add "qwer"
    ./personal-playlist playlist
