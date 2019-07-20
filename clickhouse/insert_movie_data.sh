if ! wget --spider --quiet --tries=12 --waitretry=1 --retry-connrefused "http://localhost:8123/ping" ; then
	echo >&2 'ClickHouse init process failed.'
	exit 1
fi

python3 /data/insert_movie_data.py