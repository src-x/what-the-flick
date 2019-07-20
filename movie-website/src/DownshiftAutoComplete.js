import React from "react";
import Downshift from "downshift";

const DownshiftAutoComplete =
        ({ recommendMovies, inputOnChange, onEnter, suggested_movies }) =>
         (
        <Downshift
        onChange = {item => {recommendMovies(item.title)}}
        itemToString={item => (item ? item.title : "")}
        >
        {({
        selectedItem,
        getInputProps,
        getItemProps,
        highlightedIndex,
        isOpen,
        getLabelProps
        }) => {
        return (
            <div>
                <label
                style={{ marginTop: "1rem", display: "block" }}
                {...getLabelProps()}
                >
                Choose your favourite movie
                </label>{" "}
                <br />
                <input
                {...getInputProps({
                    placeholder: "Search movies",
                    onChange: (event) => {inputOnChange(event)},
                    onKeyDown: (event, isOpen, highlightedIndex) => {onEnter(event, isOpen, highlightedIndex)}
                })}
                />
                {isOpen && suggested_movies ? (
                <div className="downshift-dropdown">
                    {suggested_movies
                    .map((item, index) => (
                        <div
                        className="dropdown-item"
                        {...getItemProps({ key: index, index, item })}
                        style={{
                            backgroundColor: highlightedIndex === index ? "lightgray" : "white",
                            fontWeight: selectedItem === item ? "bold" : "normal"
                        }}
                        >
                        {item.title}
                        </div>
                    ))}
                </div>
                ) : null}
            </div>
            )}}
        </Downshift>
        );

export default DownshiftAutoComplete;