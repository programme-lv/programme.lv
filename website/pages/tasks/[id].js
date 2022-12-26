import { useRouter } from 'next/router'

export default function Task(){
  const router = useRouter()
  const { id } = router.query

  return <p>Task: {id}</p>
}