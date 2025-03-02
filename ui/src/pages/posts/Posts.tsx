import React, {useEffect} from "react";
import {Card, CardContent, CardMedia, Chip, Container, Typography} from "@mui/material";
import Grid from '@mui/material/Grid2';
import {Post as ApiPost, PostsClient} from "../../api/posts.ts";


export interface Post {
    id: number;
    title: string;
    thumbnail: string;
    tags: string[];
}

export interface PostsProps {
    postsClient: PostsClient
}

const defaultPosts = [
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

function PostGrids({posts}: { posts: Post[] }) {
    if (posts.length === 0) {
        return <Typography>No Posts</Typography>;
    }

    return posts.map((post) => (
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
    ));
}

export default function Posts(props: PostsProps) {
    const [posts, setPosts] = React.useState<Post[]>([]);

    const mapPosts = (res: ApiPost[]): Post[] => {
        return res.map((post) => ({
            id: post.id,
            title: post.title,
            thumbnail: post.url,
            tags: post.tags,
        }));
    }

    useEffect(() => {
        props.postsClient.getPosts()
            .then((res) => {
                setPosts(mapPosts(res));
            })
            .catch((error: any) => {
                console.error(error);
                alert(error.message);
            });
    }, []);

    return (
        <Container sx={{mt: 4}}>
            <Grid container spacing={3}>
                <PostGrids posts={posts}/>
            </Grid>
        </Container>
    );
}