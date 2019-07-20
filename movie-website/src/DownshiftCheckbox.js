import React from "react";
import Downshift from "downshift";

function stateReducer(state, changes) {
  switch (changes.type) {
    case Downshift.stateChangeTypes.keyDownEnter:
    case Downshift.stateChangeTypes.clickItem:
      return {
        ...changes,
        isOpen: state.isOpen,
        highlightedIndex: state.highlightedIndex,
      }
    default:
      return changes
  }
}

const DownshiftCheckbox =
    ({ onRemoveItem, itemToString, genres, ...rest }) =>
    (
    <Downshift itemToString={item => (item ? item.genre : "")}
    {...rest}
    stateReducer ={stateReducer}
    >
          {({
            getToggleButtonProps,
            getItemProps,
            isOpen,
            selectedItem,
          }) => {
          return(
            <div className="custom-select-checkbox">
              <button {...getToggleButtonProps()} className="custom-select-button">
              <div display="flex">
                {selectedItem.length > 0
                  ? selectedItem.map(item => (
                      <button
                        className = "tag"
                        key={item.id}
                        onClick = {(e) => {onRemoveItem(e, item)}}
                      >
                        <span>x</span>
                        {' '}{item.genre}{' '}
                      </button>
                    ))
                  : 'Genres'}
                  <svg className = "arrowIcon"
                    viewBox="0 0 20 20"
                    preserveAspectRatio="none"
                    width={16}
                    fill="transparent"
                    stroke="#ebebeb"
                    strokeWidth="1.1px"
                    transform={isOpen ? 'rotate(180)' : null}
                  >
                    <path d="M1,6 L10,15 L19,6" />
                  </svg>
              <div/>
            </div>
              </button>
              {!isOpen ? null : (
                <div className= "grid-container">
                  {genres.map((item, index) => (
                    <label key = {item.id} className = "container">
                      {item.genre}
                      <input
                        type="checkbox"
                        {...getItemProps({
                          item,
                          index,
                          checked: selectedItem.includes(item)
                        })}
                        id={item.id}
                        value={item.id}
                      />
                      <span className = "checkmark">
                      </span>
                    </label>
                  ))}
                </div>
              )}
            </div>
          )}}
        </Downshift>
    );

export default DownshiftCheckbox;
