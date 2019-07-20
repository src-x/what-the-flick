import pandas as pd
from clickhouse_driver import Client

client = Client(host='localhost')

# add ratings to movies dataframe
movies_df = pd.read_csv("./ml-latest-small/movies.csv").drop(columns = 'genres')
ratings_df = pd.read_csv("./ml-latest-small/ratings.csv").drop(columns = 'timestamp')

movies_ratings_df = pd.merge(movies_df, ratings_df, on = "movieId")

# replace movieId with tmdbId
links_df = pd.read_csv("./ml-latest-small/links.csv").drop(columns = 'imdbId')
links_df['tmdbId'] = links_df['tmdbId'].astype('Int64')

movies_ratings_df = pd.merge(links_df, movies_ratings_df, on = 'movieId')
movies_ratings_df = movies_ratings_df.drop(columns= 'movieId')
movies_ratings_df = movies_ratings_df.dropna()

movies_ratings_df['tmdbId'] = movies_ratings_df['tmdbId'].astype('int64')

# Insert DF to Clickhouse
client.execute('CREATE TABLE IF NOT EXISTS test (userId Int64, tmdbId Int64, rating Float64, title String) ENGINE = Memory')
client.execute('INSERT INTO test (userId, tmdbId, rating, title) VALUES', movies_ratings_df.to_dict("records"),types_check=True)