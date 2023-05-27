import React, {useEffect, useState} from "react";
import {Button, List, ListItem, ListItemText} from '@mui/material';
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
    },[])

    const handleVote = () => {
     setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

  return (
    <List>
        {restos.map((r) => (
            <ListItem disablePadding>
                <ListItemText primary={r.name} />
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