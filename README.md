![SoundBert logo](https://raw.githubusercontent.com/HorayNarea/SoundBert/master/static/logo.png)

## SoundBert - A simple REST based soundboard in Go.
[![Build Status](https://travis-ci.org/HorayNarea/SoundBert.svg?branch=master)](https://travis-ci.org/HorayNarea/SoundBert)

----

### Installation
```sh
$ go get github.com/HorayNarea/SoundBert
$ cp config.default.toml config.toml
$ ./SoundBert
```

##### Requirements
Due to the usage of *mplayer* you have to install the **MPlayer CLI**. For further information look here:

* [Install on Ubuntu](http://www.debianadmin.com/install-mplayer-ubuntu.html)
* [Install on OSX via Homebrew](https://github.com/donmelton/MPlayerShell)
* [Install on Raspberry PI](https://rasspberrypi.wordpress.com/2012/09/02/audio-and-video-playback-on-raspberry-pi/)

### Get started
Copy **config.default.toml** to **config.toml** and adjust the properties as you like.
For example the *host* and *port* properties hold the information which ip and port the server binds to. The *sounds* property is the location for the sound files.

After doing that, simply start the server by running:
```sh
$ ./SoundBert
```

And open *http://localhost:8080/* or whatever you set in **config.toml**


### The board itself
[GET]  **/**

### List all available snippets
[GET]  **/list**

Sample response:
```javascript
{
    "foo.mp3": "foo.mp3",
    "bar/baz": "bar/baz.wav"
}
```

### Play a snippet
[POST] **/play**

Post data example:
```
filename=foo.mp3
```

Sample response:
```javascript
{"playing":"foo.mp3"}
```

### Stop playback
[GET] **/stop**

### Reload sounds
[GET] **/reload_sounds**

### Help
[GET]  **/help.html**

### License
ISC (it's like MIT but even simpler), see [LICENSE](LICENSE) for full text

### Thanks
[jlis](https://github.com/jlis) for the original [soundbeard](https://github.com/jlis/soundbeard) in NodeJS
