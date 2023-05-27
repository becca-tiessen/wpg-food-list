import * as React from 'react';
import { 
    Button, 
    Dialog, 
    DialogActions, 
    DialogContent, 
    DialogTitle, 
    Rating,
    List,
    ListItem,
    ListItemText,
} from '@mui/material';

export default function VoteDialog(props) {
  const { onClose, open, factors } = props;

  const handleClose = () => {
    onClose();
  };

  const handleSubmit = () => {
    alert('submitted rating!');
    handleClose();
  }

  return (
    <Dialog onClose={handleClose} open={open}>
      <DialogTitle>Vote for restaurant</DialogTitle>
      <DialogContent>
        <List>
            {factors.map(f => (
                <ListItem>
                    <ListItemText>{f}</ListItemText>
                    <Rating name="half-rating" value={null}/>
                </ListItem>
            ))}
        </List>
      </DialogContent>
      <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleSubmit}>Submit rating</Button>
        </DialogActions>
    </Dialog>
  );
}