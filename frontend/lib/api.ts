import fs from 'fs'
import { join } from 'path'

const postsDirectory = join(process.cwd(), '_posts')

export function getPostSlugs() {
  return fs.readdirSync(postsDirectory)
}

export async function getPostBySlug(slug: string) {

  try {
    const res = await fetch('http://127.0.0.1:4000/article/'+slug);
    const data = await res.json();

    const convertedData ={
      oid: data.oid,
      slug: data.slug,
      title: data.title,
      excerpt: data.excerpt,
      date: data.date,
      coverImage: data.coverImage,
      content: data.content,
      author: {
        name: data.author_name,
        picture: data.author_picture_url
      },
      ogImage: {url: data.open_graph_image_url}
    }
 
    return convertedData
  } catch (error) {
    console.log(error);
  }
}

export async function getAllPosts(fields: string[] = []) {

  try {
    const res = await fetch('http://127.0.0.1:4000/article');    
    const data = await res.json();

    const convertedData = data
      .map(p => ({
        oid: p.oid,
        slug: p.slug,
        title: p.title,
        excerpt: p.excerpt,
        date: p.date,
        coverImage: p.coverImage,
        author: {
          name: p.author_name,
          picture: p.author_picture_url
        }
      }))
   
    return convertedData
  } catch (err) {
    console.log(err);
  }

   
}
