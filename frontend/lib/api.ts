import fs from 'fs'
import { join } from 'path'
import matter from 'gray-matter'

const postsDirectory = join(process.cwd(), '_posts')

export function getPostSlugs() {
  return fs.readdirSync(postsDirectory)
}

export function getPostBySlug(slug: string, fields: string[] = []) {
  const realSlug = slug.replace(/\.md$/, '')
  const fullPath = join(postsDirectory, `${realSlug}.md`)
  const fileContents = fs.readFileSync(fullPath, 'utf8')
  const { data, content } = matter(fileContents)

  type Items = {
    [key: string]: string
  }
  
  const items: Items = {}

  // Ensure only the minimal needed data is exposed
  fields.forEach((field) => {
    if (field === 'slug') {
      items[field] = realSlug
    }
    if (field === 'content') {
      items[field] = content
    }

    if (typeof data[field] !== 'undefined') {
      items[field] = data[field]
    }
  })

  return items
}

export async function getAllPosts(fields: string[] = []) {

  try {
    const res = await fetch('http://127.0.0.1:4000/article');
    
    const data = await res.json();
    console.log(data)
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
