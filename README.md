# tmdb
CLI to search movies and rename files based in TheMovieDB database.

## Installation
### MAC
```
brew tap ssuareza/brew git@github.com:ssuareza/homebrew-brew
brew install ssuareza/brew/gssh -f
```

### Linux
```
curl -sLo tmdb https://github.com/ssuareza/tmdb/releases/download/v0.0.1/tmdb-v0.0.1-linux-amd64
chmod +x tmdb
sudo mv tmdb /usr/local/bin/
```

## Examples
### Search movie
```tmbd search Avengers```

### Rename movie based in TheMovieDB database
```tmbd rename Avengers.Endgame.2019.BlueRay.mkv```

### Rename movie and move file to another directory
```tmdb rename /path/Avengers.Endgame.2019.BlueRay/Avengers.Endgame.2019.BlueRay.mkv --move /path/destination/```