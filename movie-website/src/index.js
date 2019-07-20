import React, { Component } from "react";
import ReactDOM from "react-dom";
import axios from "axios";
import DownshiftCheckbox from "./DownshiftCheckbox"
import DownshiftAutoComplete from "./DownshiftAutoComplete";
import "./styles.css";

const genres = [
  {id: 1, genre:'Animation'},
  {id: 2, genre:'Romance'},
  {id: 3, genre:'War'},
  {id: 4, genre:'Comedy'},
  {id: 5, genre:'Film-Noir'},
  {id: 6, genre:'Crime'},
  {id: 7, genre:'Sci-Fi'},
  {id: 8, genre:'Fantasy'},
  {id: 9, genre:'Children'},
  {id: 10, genre:'Western'},
  {id: 11, genre:'Horror'},
  {id: 12, genre:'Thriller'},
  {id: 13, genre:'Mystery'},
  {id: 14, genre:'Drama'},
  {id: 15, genre:'Documentary'},
  {id: 16, genre:'Musical'},
  {id: 17, genre:'Adventure'},
  {id: 18, genre:'Action'},
  {id: 19, genre:'(no genres listed)'},
  {id: 20, genre:'IMAX'}
];

class Main extends Component {
  constructor(props) {
    super(props);
    this.state = {
      suggested_movies: [],
      searched_movies: [],
      searched_movie_posters: [],
      recommended_movies: [],
      recommended_movie_posters: [],
      selected_items: []
    };

    this.suggestMovies = this.suggestMovies.bind(this);
    this.searchMovies = this.searchMovies.bind(this);
    this.recommendMovies = this.recommendMovies.bind(this);
    this.inputOnChange = this.inputOnChange.bind(this);
  }
  
  add_remove = (selectedItem) => {
    if (this.state.selected_items.includes(selectedItem)) {
      this.removeItem(selectedItem);
    } else {
      this.addSelectedItem(selectedItem);
    }
  }

  removeItem(item) {
    this.setState(({ selected_items }) => {
      return {
        selected_items: selected_items.filter(i => i !== item)
      };
    });
  }

  addSelectedItem(item) {
    this.setState(
      ({ selected_items }) => ({
        selected_items: [...selected_items, item]
      })
    );
  }

  onItemChanged = (item) => {
    const copy = [...this.state.selected_items]
    copy.splice(item.index, 1, item.value)
    this.setState({ selected_items: copy })
  }

  itemToString(i) {
    return i ? i.name : ''
  }

  handleChange = selectedItems => {
    this.setState({selected_items: selectedItems})
  }

  addItem = (item, callbackFunction) => {
    this.setState(
      ({ selected_items }) => ({
        selected_items: [...selected_items, item]
      }),
      callbackFunction
    );
  };

  inputOnChange(event) {
    if (!event.target.value) {
      this.setState({ suggested_movies: [] })
      return;
    }
    this.suggestMovies(event.target.value);
  }

  suggestMovies(movie) {
    const moviesURL = `http://localhost:8090/suggest/${movie}`;
    axios.get(moviesURL).then(response => {
      this.setState({ suggested_movies: response.data.Movie });
    });
  }

  searchMovies(movie) {
    const moviesURL = `http://localhost:8090/search/${movie}`;
    axios.get(moviesURL).then(response => {
      const axeids = []
      this.setState({ searched_movies: response.data.Movie });
      if (response.data.Movie != null) {
        for (let i = 0; i < response.data.Movie.length; i++) {
        let axeid = axios.get(`https://api.themoviedb.org/3/movie/${response.data.Movie[i].id}?api_key=`)
        .then(response => response.data);
        axeids.push(axeid);
        }
        axios.all(axeids).then(axios.spread((...res) => this.setState({searched_movie_posters : res})))
      }
    });
  }

  recommendMovies(movie) {
    const movie_text = movie.replace(/'/g, "''");
    const moviesURL = `http://localhost:8091/recommend/${movie_text}`;
    axios.get(moviesURL).then(response => {
      const axeids = [];
      this.setState({ recommended_movies: response.data.Movie });
      if (response.data.Movie != null) {
      for (let i = 0; i < response.data.Movie.length; i++) {
      let axeid = axios.get(`https://api.themoviedb.org/3/movie/${response.data.Movie[i].id}?api_key=`)
      .then(response => response.data);
      axeids.push(axeid);
      }
      axios.all(axeids).then(axios.spread((...res) => this.setState({recommended_movie_posters : res})))
    }
  })
  }

  handleSubmit = (e) => {
    e.preventDefault();
    console.log(e.target.value);
  }

  onRemoveItem = (e, item) => {
    e.preventDefault();
    e.stopPropagation();
    this.removeItem(item);
  }
  onEnter = (event, isOpen, highlightedIndex) => {
    if (event.key === 'Enter') {
        this.setState({ suggested_movies: [] })
        if (!isOpen || highlightedIndex === null) {
        event.preventDownshiftDefault = true;
        if (this.state.selected_items.length !== 0) {
        const genre_string = this.state.selected_items.map((item, index) => (
            ` ${item.genre} `
        )
        ).join('|')
        this.searchMovies(event.target.value.trim() + " @genres: {" + genre_string + "}")
        }
        else {
            this.searchMovies(event.target.value.trim())
        }
        }
    }
    } 
  getRemoveButtonProps = ({onClick} = {}) => {
    return {
      onClick: e => {
        onClick && onClick(e)
        e.stopPropagation()
      },
    }
  }
  render() {
  const POSTER_PATH = 'http://image.tmdb.org/t/p/w154';
    return (
      <div className="row">
      <div className="left">
      <h2>Menu</h2>
      <DownshiftCheckbox
          selectedItem={this.state.selected_items}
          onChange={this.add_remove}
          onItemChanged={this.onItemChanged}
          onRemoveItem={this.onRemoveItem}
          genres= {genres}
          itemToString={this.itemToString}
        />
      <DownshiftAutoComplete
        recommendMovies = {this.recommendMovies}
        suggested_movies = {this.state.suggested_movies}
        inputOnChange = {this.inputOnChange}
        onEnter = {this.onEnter}
        >
      </DownshiftAutoComplete>
      <br/>
      <div>        
      {this.state.searched_movies != null && this.state.searched_movies.length > 0 ? ( 
      <div>Did you mean?</div>
      ): null}
      {this.state.searched_movies ? (
    <div className="leftGrid">        
    <div className = "leftWrapper">
      {this.state.searched_movie_posters.map(movie => (
        <a href="#/" onClick = {() => this.recommendMovies(movie.title + ` (${movie.release_date.substring(0,4)})`)} key = {movie.id}>
          <img src={`${POSTER_PATH}${movie.poster_path}`} alt = "movie poster"/>
        <h3><span><b>{movie.title}</b></span></h3>
        </a>
      ))}
      </div>
      </div>
      ) : <div> sorry, not found</div>}
      </div>
      </div>
      <div className="right">
    <h2>Movie Recommendations</h2>
    <div className="rightGrid">        
    <div className = "rightWrapper">
      {this.state.recommended_movie_posters.map(movie => (
        <a href="#/" key = {movie.id}>
          <p className = "moviepTag">
            <img src={`${POSTER_PATH}${movie.poster_path}`} alt = "movie poster"/> 
          <span className = "list">
            <li>Rating: {movie.vote_average}</li>
            <li>Runtime: {movie.runtime}</li> 
            <li>Genres: {movie.genres.map(genre => (<li className = "subList" key = {genre.id}>{genre.name}</li>))}</li>
            {/* <li>Companies:{movie.production_companies.map(company => (<li className = "subList">{company.name}</li>))}</li>
            <li>Budget: {movie.budget}</li> 
            <li>Revenue: {movie.revenue}</li>  */}
          </span> 
          <span>Release Date: {movie.release_date}</span> 
          </p>
        <span><h3><b>{movie.title}</b></h3></span>
        <p className = "moviepTag">Description:</p>
        <p>{movie.overview}</p>
        </a>
      ))}
      </div>
      </div>
      </div>
  </div>
    );
  }
}

export default Main;

const rootElement = document.getElementById("root");
ReactDOM.render(<Main />, rootElement);
