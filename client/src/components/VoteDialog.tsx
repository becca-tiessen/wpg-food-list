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

export interface VoteDialogProps {
  open: boolean;
  selectedValue: string;
  onClose: (value: string) => void;
  factors: Array<string>;
}

export default function VoteDialog(props: VoteDialogProps) {
  const { onClose, selectedValue, open, factors } = props;

  const handleClose = () => {
    onClose(selectedValue);
  };

  return (
    <Dialog onClose={handleClose} open={open}>
      <DialogTitle>Vote for restaurant</DialogTitle>
      <DialogContent>
        <List>
            {factors.map(f => (
                <ListItem>
                    <ListItemText>{f}</ListItemText>
                    <Rating name="half-rating" value={4.5} precision={0.5} />
                </ListItem>
            ))}
        </List>
      </DialogContent>
      <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleClose}>Submit rating</Button>
        </DialogActions>
    </Dialog>
  );
}