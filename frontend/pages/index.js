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

  if (error) {
    return <div>Failed to load {error.toString()}</div>
  }

  if (!data) {
    return <div>Loading...</div>
  }

  return (
    <div className={styles.container}>
      <p>Hello World</p>
      <p>data : {data.message}</p>
    </div>
  )
}
