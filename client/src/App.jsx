import React from "react";
import {Container, ThemeProvider, createTheme} from '@mui/material';
import { amber, deepOrange } from '@mui/material/colors';
import NavMenu from './components/NavMenu';
import FoodList from './components/FoodList';

// spend more time on this eventually
const theme = createTheme({
    palette: {
      primary: {
        main: "#f6856a"
        // #f6856a
        //#fa94b5
      },
      secondary: {
        main: deepOrange[500]
      }
    },
  });

const App = () => {
    return (
        <>
        <ThemeProvider theme={theme}>
        <NavMenu>
            <Container maxWidth="md">
            <h1>
                welcome to winnipeg's food list.
            </h1>
            <FoodList/>
            </Container>
        </NavMenu>
        </ThemeProvider>;
        </>
    )
}

export default App