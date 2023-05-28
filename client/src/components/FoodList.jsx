import React, {useEffect, useState} from "react";
import {Box, Button, Divider, List, ListItem, ListItemText, Rating, Stack} from '@mui/material';
import VoteDialog from "./VoteDialog";

export default function FoodList() {
    const [restos, setRestos] = useState([]);
    const factors = ['value', 'vibes', 'flavour']
    const [open, setOpen] = useState(false);

    const getRestos = async () => {
        await fetch('/restaurants')
        .then(res => res.json())
        .then(
        (result) => {
            setRestos(result.restaurants);
        }).catch((err) => {
            console.log(err)
        })
    };

    useEffect(() => {
        getRestos();
    },[]);

    const handleVote = () => {
     setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const renderRestoRatings=()=> {
        return (
            //this will need to be its own component eventually
            <Stack justifyContent="center" spacing={3} direction="row" divider={<Divider orientation="vertical" flexItem />}>
                <Box>value: <Rating size="small" value={5} readOnly/></Box>
                <Box>vibes: <Rating precision={0.5} size="small" value={3.5} readOnly/></Box>
                <Box>flaves: <Rating size="small" value={4} readOnly/></Box>
            </Stack>
        )
    }

  return (
    <List>
        {restos.map((r) => (
            <ListItem>
                <ListItemText 
                primary={r.name} 
                secondary={renderRestoRatings()}
                />
                <Button variant="outlined" onClick={handleVote}>
                    Vote
                </Button>
                <VoteDialog
                    open={open}
                    onClose={handleClose}
                    factors={factors}
                />
            </ListItem>
        ))}
    </List>
  )
}