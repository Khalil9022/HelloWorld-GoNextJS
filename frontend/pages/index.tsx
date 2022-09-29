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
      setData(data)
    } catch (error) {
      setError(error);
    }
  }

  return (
    <div className='container mx-auto'>
      <h1 className='text-sky-400'>Selamat Datang Di websiteku</h1>
      <br></br>
      {error && <div>Failed to load{error.toString()}</div>}
      {
        !data ? <div>Loading...</div>
          : (
            (data?.data ?? []).length === 0
          )
      }

      <Input onSuccess={getData} />
      {data?.data && data?.data?.map((item, index) => (
        <div key={index}>
          <span>ID: {item.ID} task: {item.task} </span>
          <input type="checkbox" defaultChecked={item.done} />
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
      text: formData.get("data")
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
        <input name='data' type="text" />
        <button className='rounded bg-sky-500 '>Submit</button>
      </form>
    </div>
  )
}
