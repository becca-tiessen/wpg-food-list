import React from "react";
import {Container} from '@mui/material';
import NavMenu from './components/NavMenu';
import FoodList from './components/FoodList';

const App = () => {
    return (
        <>
        <NavMenu>
            <Container maxWidth="sm">
            <h1>
                welcome to winnipeg's food list.
            </h1>
            <FoodList/>
            </Container>
        </NavMenu>
        </>
    )
}

export default App