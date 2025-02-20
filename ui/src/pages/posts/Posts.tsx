import React from "react";
import {Card, CardContent, CardMedia, Chip, Container, Typography} from "@mui/material";
import Grid from '@mui/material/Grid2';

const postsData = [
    {
        id: 1,
        title: "React Tutorial",
        thumbnail: "https://via.placeholder.com/320x180",
        tags: ["React", "JavaScript", "Frontend"],
    },
    {
        id: 2,
        title: "Go Fiber Guide",
        thumbnail: "https://via.placeholder.com/320x180",
        tags: ["Go", "Backend", "Fiber"],
    },
    {
        id: 3,
        title: "Docker Basics",
        thumbnail: "https://via.placeholder.com/320x180",
        tags: ["Docker", "DevOps", "Containers"],
    },
];

export default function Posts() {
    return (
        <Container sx={{mt: 4}}>
            <Grid container spacing={3}>
                {postsData.map((post) => (
                    <Grid key={post.id} size={{xs: 12, sm: 6, md: 4}}>
                        <Card>
                            <CardMedia component="img" height="180" image={post.thumbnail} alt={post.title}/>
                            <CardContent>
                                <Typography variant="h6">{post.title}</Typography>
                                <Grid container spacing={1} sx={{mt: 1}}>
                                    {post.tags.map((tag, index) => (
                                        <Grid key={index}>
                                            <Chip label={tag} size="small"/>
                                        </Grid>
                                    ))}
                                </Grid>
                            </CardContent>
                        </Card>
                    </Grid>
                ))}
            </Grid>
        </Container>
    );
};
