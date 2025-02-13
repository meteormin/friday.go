import * as React from 'react';
import Avatar from '@mui/material/Avatar';
import AvatarGroup from '@mui/material/AvatarGroup';
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Chip from '@mui/material/Chip';
import Grid from '@mui/material/Grid2';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import FormControl from '@mui/material/FormControl';
import InputAdornment from '@mui/material/InputAdornment';
import OutlinedInput from '@mui/material/OutlinedInput';
import {styled} from '@mui/material/styles';
import SearchRoundedIcon from '@mui/icons-material/SearchRounded';
import RssFeedRoundedIcon from '@mui/icons-material/RssFeedRounded';

interface CardData {
    img: string;
    tag: string;
    title: string;
    description: string;
    author: { name: string; avatar: string };
}

const cardData = [
    {
        img: 'https://picsum.photos/800/450?random=3',
        tag: 'Design',
        title: 'Designing for the future: trends and insights',
        description:
            'Stay ahead of the curve with the latest design trends and insights. Our design team shares their expertise on creating intuitive and visually stunning user experiences.',
        author: {name: 'Kate Morrison', avatar: '/static/images/avatar/7.jpg'},
    },
    {
        img: 'https://picsum.photos/800/450?random=4',
        tag: 'Company',
        title: "Our company's journey: milestones and achievements",
        description:
            "Take a look at our company's journey and the milestones we've achieved along the way. From humble beginnings to industry leader, discover our story of growth and success.",
        author: {name: 'Cindy Baker', avatar: '/static/images/avatar/3.jpg'},
    },
    {
        img: 'https://picsum.photos/800/450?random=6',
        tag: 'Product',
        title: 'Maximizing efficiency with our latest product updates',
        description:
            'Our recent product updates are designed to help you maximize efficiency and achieve more. Get a detailed overview of the new features and improvements that can elevate your workflow.',
        author: {name: 'Travis Howard', avatar: '/static/images/avatar/2.jpg'},
    },
];

const StyledCard = styled(Card)(({theme}) => ({
    display: 'flex',
    flexDirection: 'column',
    padding: 0,
    height: '100%',
    backgroundColor: (theme).palette.background.paper,
    '&:hover': {
        backgroundColor: 'transparent',
        cursor: 'pointer',
    },
    '&:focus-visible': {
        outline: '3px solid',
        outlineColor: 'hsla(210, 98%, 48%, 0.5)',
        outlineOffset: '2px',
    },
}));

const StyledCardContent = styled(CardContent)({
    display: 'flex',
    flexDirection: 'column',
    gap: 4,
    padding: 16,
    flexGrow: 1,
    '&:last-child': {
        paddingBottom: 16,
    },
});

const StyledTypography = styled(Typography)({
    display: '-webkit-box',
    WebkitBoxOrient: 'vertical',
    WebkitLineClamp: 2,
    overflow: 'hidden',
    textOverflow: 'ellipsis',
});

function Author({authors}: { authors: { name: string; avatar: string }[] }) {
    return (
        <Box
            sx={{
                display: 'flex',
                flexDirection: 'row',
                gap: 2,
                alignItems: 'center',
                justifyContent: 'space-between',
                padding: '16px',
            }}
        >
            <Box
                sx={{display: 'flex', flexDirection: 'row', gap: 1, alignItems: 'center'}}
            >
                <AvatarGroup max={3}>
                    {authors.map((author, index) => (
                        <Avatar
                            key={index}
                            alt={author.name}
                            src={author.avatar}
                            sx={{width: 24, height: 24}}
                        />
                    ))}
                </AvatarGroup>
                <Typography variant="caption">
                    {authors.map((author) => author.name).join(', ')}
                </Typography>
            </Box>
            <Typography variant="caption">July 14, 2021</Typography>
        </Box>
    );
}

export function Search() {
    return (
        <FormControl sx={{width: {xs: '100%', md: '25ch'}}} variant="outlined">
            <OutlinedInput
                size="small"
                id="search"
                placeholder="Search…"
                sx={{flexGrow: 1}}
                startAdornment={
                    <InputAdornment position="start" sx={{color: 'text.primary'}}>
                        <SearchRoundedIcon fontSize="small"/>
                    </InputAdornment>
                }
                inputProps={{
                    'aria-label': 'search',
                }}
            />
        </FormControl>
    );
}

interface Tag {
    id: number;
    name: string;
    selected: boolean;
}


interface CardGridProps {
    card: CardData;
    handleFocus: (index: number) => void;
    handleBlur: () => void;
}

function CartGrid(props: CardGridProps) {
    return (
        <Grid size={{xs: 12, md: 4}}>
            <StyledCard
                variant="outlined"
                onFocus={() => props.handleFocus(2)}
                onBlur={props.handleBlur}
                tabIndex={0}
                className={'Mui-focused'}
                sx={{height: '100%'}}
            >
                <CardMedia
                    component="img"
                    alt="green iguana"
                    image={props.card.img}
                    sx={{
                        height: {sm: 'auto', md: '50%'},
                        aspectRatio: {sm: '16 / 9', md: ''},
                    }}
                />
                <StyledCardContent>
                    <Typography gutterBottom variant="caption" component="div">
                        {props.card.tag}
                    </Typography>
                    <Typography gutterBottom variant="h6" component="div">
                        {props.card.title}
                    </Typography>
                    <StyledTypography variant="body2" color="text.secondary" gutterBottom>
                        {props.card.description}
                    </StyledTypography>
                </StyledCardContent>
                <Author authors={[props.card.author]}/>
            </StyledCard>
        </Grid>
    );
}

export default function Posts() {
    const [focusedCardIndex, setFocusedCardIndex] = React.useState<number | null>(
        null,
    );

    const [tags, setTags] = React.useState<Tag[]>([
        {id: 1, name: 'All', selected: true},
        {id: 2, name: 'Engineering', selected: false},
        {id: 3, name: 'Product', selected: false},
    ]);

    const handleFocus = (index: number) => {
        setFocusedCardIndex(index);
    };

    const handleBlur = () => {
        setFocusedCardIndex(null);
    };

    const handleClick = (id: number) => {
        console.info('You clicked the filter chip.');

        setTags(prevState =>
            prevState.map(tag => {
                tag.selected = tag.id === id;
                return tag;
            })
        );
    };

    return (
        <Box sx={{display: 'flex', flexDirection: 'column', gap: 4}}>
            <div>
                <Typography variant="h1" gutterBottom>
                    Posts
                </Typography>
                <Typography>Stay in the loop with the latest about our products</Typography>
            </div>
            <Box
                sx={{
                    display: {xs: 'flex', sm: 'none'},
                    flexDirection: 'row',
                    gap: 1,
                    width: {xs: '100%', md: 'fit-content'},
                    overflow: 'auto',
                }}
            >
                <Search/>
                <IconButton size="small" aria-label="RSS feed">
                    <RssFeedRoundedIcon/>
                </IconButton>
            </Box>
            <Box
                sx={{
                    display: 'flex',
                    flexDirection: {xs: 'column-reverse', md: 'row'},
                    width: '100%',
                    justifyContent: 'space-between',
                    alignItems: {xs: 'start', md: 'center'},
                    gap: 4,
                    overflow: 'auto',
                }}
            >
                <Box
                    sx={{
                        display: 'inline-flex',
                        flexDirection: 'row',
                        gap: 3,
                        overflow: 'auto',
                    }}
                >
                    {tags.map(tag => (
                        <Chip
                            key={tag.id}
                            onClick={() => handleClick(tag.id)}
                            size="medium"
                            label={tag.name}
                            sx={tag.selected ? {} : {
                                backgroundColor: 'transparent',
                                border: 'none',
                            }}
                        />
                    ))}
                </Box>
                <Box
                    sx={{
                        display: {xs: 'none', sm: 'flex'},
                        flexDirection: 'row',
                        gap: 1,
                        width: {xs: '100%', md: 'fit-content'},
                        overflow: 'auto',
                    }}
                >
                    <Search/>
                    <IconButton size="small" aria-label="RSS feed">
                        <RssFeedRoundedIcon/>
                    </IconButton>
                </Box>
            </Box>
            <Grid container spacing={2} columns={12}>
                {cardData.map(card => (
                    <CartGrid card={card} handleFocus={handleFocus} handleBlur={handleBlur}/>
                ))}
            </Grid>
        </Box>
    );
}