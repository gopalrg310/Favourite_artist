# Favourite_artist
To my favourite artist

Please run the code in one terminal by running command go run main.go

Hit the API is

GET
http://localhost:8080/top-track?region=india

Response FORMAT

`{
    "artist": {
        "name": "Coldplay",
        "info": "https://upload.wikimedia.org/wikipedia/commons/2/2e/ColdplayBBC071221_%28cropped%29.jpg"
    },
    "name": "Yellow",
    "lyrics": {
        "lyrics_body": "Look at the stars\nLook how they shine for you\n\nAnd everything you do\n\nYeah, they were all yellow\n\nI came along\nI wrote a song for you\nAnd all the things you do\n\nAnd it was called Yellow\n\nSo then I took my turn\nOh, what a thing to have done\nAnd it was all yellow\n...\n\n******* This Lyrics is NOT for Commercial use *******"
    }
}`


To Run test cases
`go test -v`


And do 
`go build`
