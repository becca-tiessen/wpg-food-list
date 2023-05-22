import React from "react";
import {Box, Container, List, ListItem, ListItemText} from '@mui/material';

const App = () =>{
    return (
        <>
        <Container maxWidth="sm">
        <h1>
            welcome to winnipeg's food list.
        </h1>
        <List>
          <ListItem disablePadding>
              <ListItemText primary="Sous Sol" />
          </ListItem>
          <ListItem disablePadding>
              <ListItemText primary="Deer + Almond" />
          </ListItem>
        </List>
        </Container>
        </>
    )
}

export default App