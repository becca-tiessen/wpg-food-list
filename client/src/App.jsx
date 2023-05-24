import React from "react";
import {Button, Container, List, ListItem, ListItemText} from '@mui/material';
import NavMenu from './components/NavMenu.tsx';
import VoteDialog from './components/VoteDialog.tsx';

const App = () => {
    const factors = ['value', 'vibes', 'flavour']
    const [open, setOpen] = React.useState(false);

    const handleVote = () => {
     setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

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
                <Button variant="outlined" onClick={handleVote}>
                    Vote
                </Button>
                <VoteDialog
                    open={open}
                    onClose={handleClose}
                    factors={factors}
                />
            </ListItem>
            <ListItem disablePadding>
                <ListItemText primary="Sous Sol" />
                <Button variant="outlined" onClick={handleVote}>
                    Vote
                </Button>
                <VoteDialog
                    open={open}
                    onClose={handleClose}
                    factors={factors}
                />
            </ListItem>
            </List>
            
            </Container>
        </NavMenu>
        </>
    )
}

export default App