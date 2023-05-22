import React from "react";
import {Container, List, ListItem, ListItemText, Rating} from '@mui/material';
import NavMenu from './components/NavMenu.tsx'

const App = () =>{
    return (
        <>
        <NavMenu>
            <Container maxWidth="sm">
            <h1>
                welcome to winnipeg's food list.
            </h1>
            <List>
            <ListItem disablePadding>
                <ListItemText primary="Deer + Almond" />
                <Rating name="half-rating" value={5.0} precision={0.5} />
            </ListItem>
            <ListItem disablePadding>
                <ListItemText primary="Sous Sol" />
                <Rating name="half-rating" value={4.5} precision={0.5} />
            </ListItem>
            </List>
            </Container>
        </NavMenu>
        </>
    )
}

export default App