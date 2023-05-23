import React from "react";
import {Button, Container, List, ListItem, ListItemText, Rating} from '@mui/material';
import NavMenu from './components/NavMenu.tsx';
import VoteDialog from './components/VoteDialog.tsx';

const emails = ['username@gmail.com', 'user02@gmail.com'];

const App = () => {
    const factors = ['value', 'vibes', 'flavour']
    const [open, setOpen] = React.useState(false);
    const [selectedValue, setSelectedValue] = React.useState(emails[1]);

    const handleVote = () => {
     setOpen(true);
    };

    const handleClose = (value) => {
        setOpen(false);
        setSelectedValue(value);
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
                    selectedValue={selectedValue}
                    open={open}
                    onClose={handleClose}
                    factors={factors}
                />
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