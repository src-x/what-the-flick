FROM yandex/clickhouse-server
RUN apt-get update
RUN apt-get install python3 -y
RUN apt-get install python3-pip -y
RUN apt-get install unzip -y
RUN pip3 install pandas
RUN pip3 install clickhouse_driver
ADD . /data
WORKDIR /data
RUN wget http://files.grouplens.org/datasets/movielens/ml-latest-small.zip
RUN unzip ml-latest-small.zip
RUN rm ml-latest-small.zip
RUN sed -i '/exec $gosu/i\    \/data\/insert_movie_data.sh \&' /entrypoint.sh