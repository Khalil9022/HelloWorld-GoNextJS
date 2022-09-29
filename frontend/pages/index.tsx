import { useState, useEffect } from 'react'
import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'


export default function Home() {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null)

  useEffect(() => {
    getData()
  }, []);

  const getData = async () => {
    try {
      const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL)
      const data = await res.json()
      console.log(data);

      setData(data)
    } catch (error) {
      setError(error);
    }
  }

  const handleUpdate = async (id) => {
    try {
      await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/update/${id}`, {
        method: "PUT"
      })
      getData()
    } catch (error) {
      setError(error);
    }
  }

  const handleDelete = async (id) => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/delete/${id}`, {
        method: "DELETE"
      })
      getData()
    } catch (error) {
      setError(error);
    }
  }

  return (
    <div className='container mx-auto '>
      <br></br>
      <p className='text-sky-400 text-2xl font-serif'>Todolist</p>
      <br></br>

      {error && <div>Failed to load{error.toString()}</div>}
      {
        !data ? <div>Loading...</div>
          : (
            (data?.data ?? []).length === 0
          )
      }
      <Input onSuccess={getData} />
      <br></br>
      <p className='text-dark-200 text-xl'>List yang mau dikerjakan :</p>
      <hr />
      <br />
      {data?.data && data?.data?.map((item, index) => (
        <div key={index}>
          <input type="checkbox" checked={item.done} />
          <span>ID: {item.ID} task: {item.task} </span>
          <button onClick={() => handleUpdate(item.ID)} className='bg-sky-500 rounded' style={{ padding: 5, margin: 5, color: 'white' }} >Done</button>
          <button onClick={() => handleDelete(item.ID)} className='bg-red-600 rounded' style={{ padding: 5, margin: 5, color: 'white' }}>Delete</button>
          <br></br>
        </div>
      ))}

      {/* {data.data.map((item, index) => (
        <p key={index}>{item}</p>
      ))} */}
    </div>
  )
}

function Input({ onSuccess }) {
  const [data, setData] = useState(null)
  const [error, setError] = useState(null)

  const handleSubmit = async (e) => {
    e.preventDefault()
    const formData = new FormData(e.currentTarget)
    const body = {
      task: formData.get("data")
    }

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/send`, {
        method: 'POST',
        body: JSON.stringify(body)
      })
      const data = await res.json()
      setData(data.message)
      onSuccess()
    } catch (error) {
      setError(error)
    }
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input className='rounded border-slate-200 placeholder-slate-400 contrast-more:border-slate-400 contrast-more:placeholder-slate-500"' name='data' type="text" style={{ padding: 3 }} />
        <button className='rounded bg-sky-500 ' style={{ padding: 3, margin: 5, color: 'white' }} >Submit</button>
      </form>
    </div>
  )
}
