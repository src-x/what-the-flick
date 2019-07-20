import pandas as pd
import redis
from redis import ConnectionError
from redisearch import Client, TextField, AutoCompleter, Suggestion, TagField
import time

# movie dataframe
links_df = pd.read_csv("./ml-latest-small/links.csv").drop(columns = 'imdbId')
links_df['tmdbId'] = links_df['tmdbId'].astype('Int64')
movie_df = pd.read_csv("./ml-latest-small/movies.csv")

movie_df = pd.merge(links_df, movie_df, on = 'movieId')
movie_df = movie_df.drop(columns= 'movieId')
movie_df = movie_df.dropna()

movie_df['tmdbId'] = movie_df['tmdbId'].astype('int64')
movie_df['tmdbId'] = movie_df['tmdbId'].astype(str)
movie_df = movie_df.drop_duplicates(subset='tmdbId', keep='first')
movie_df = movie_df.reset_index(drop=True)

def insert():
    # insertion of search/suggestion data
    suggestion_client = Client('movie')
    suggestion_client.create_index([TextField('title'), TagField('genres', separator = '|')])

    for i in range(0, len(movie_df)):
        suggestion_client.add_document(movie_df['tmdbId'][i], title = movie_df['title'][i], genres = movie_df['genres'][i])

    # insertion of auto-completion data
    completion_client = AutoCompleter('ac')

    for i in range(0, len(movie_df)):
        completion_client.add_suggestions(Suggestion(movie_df['title'][i]))

def connect_insert(ip_port=6379, timeout=60):

    t = time.time() + timeout
    rs = redis.Redis(host='localhost', port=ip_port, db=0)

    while time.time() < t:
        try:
            rs.ping()
            print('redis is ready: ' + repr(ip_port))
            insert()
            return

        except redis.ConnectionError as e:
            print('can not connect to redis: ' + repr(ip_port) + ' ' + repr(e))
            time.sleep(0.1)
            continue
    else:
        print('cannot connect to redis: ' + repr(ip_port))
        raise

connect_insert()