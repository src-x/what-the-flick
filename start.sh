cd ./redisearch
git clone https://github.com/RediSearch/RediSearch.git
rm ./Redisearch/Dockerfile
mv ./RediSearch/* ./
rm -rf ./RediSearch
cd ../
# API_KEY = d2e35f816b3f67cf44a23f85a425b0a9
echo ENTER YOUR API KEY FROM TMDB
read -sp 'api_key: ' API_KEY
echo
sed -i '' "s/api_key=/&$API_KEY/" ./movie-website/src/index.js 
docker-compose up