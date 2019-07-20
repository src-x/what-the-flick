FROM redislabsmodules/rmbuilder:latest as builder

# Build the source
ARG GIT_DESCRIBE_VERSION
ADD . /RSBUILD
WORKDIR /RSBUILD
RUN set -ex;\
    ./srcutil/get_gtest.sh; \
    rm -rf docker-build; \
    mkdir docker-build; \
    cd docker-build; \
    cmake -DCMAKE_BUILD_TYPE=RelWithDebInfo ../ -DGIT_DESCRIBE_VERSION="$GIT_DESCRIBE_VERSION" -DRS_RUN_TESTS=ON; \
    make -j4; \
    pip install git+https://github.com/RedisLabsModules/RLTest; \
    pip install redis-py-cluster; \
    ctest -V;

# Package the runner
FROM redis:latest
ENV LIBDIR /usr/lib/redis/modules
WORKDIR /data
RUN set -ex;\
    mkdir -p "$LIBDIR";

COPY --from=builder /RSBUILD/docker-build/redisearch.so  "$LIBDIR"

RUN apt-get update && apt-get install wget unzip python3 python3-pip -y
RUN pip3 install pandas redisearch redis
ADD . /movie_data
WORKDIR /movie_data
RUN wget http://files.grouplens.org/datasets/movielens/ml-latest-small.zip\
&& unzip ./ml-latest-small.zip\
&& rm ./ml-latest-small.zip
RUN sed -i '/args.append(Field.NOINDEX)/a \\n        Field.__init__(self, name, *args)' /usr/local/lib/python3.7/dist-packages/redisearch/client.py\
&& sed -i '/Field.__init__(self, name, Field.TAG)/,/args.append(separator)/c\        args = [Field.TAG, Field.SEPARATOR, separator]' /usr/local/lib/python3.7/dist-packages/redisearch/client.py
CMD ["sh", "-c", "python3 /movie_data/suggestion_completion_data.py & redis-server --loadmodule /usr/lib/redis/modules/redisearch.so"]