import Sidebar from '../components/sidebar'

export default function Home() {
  return (
    <div class="container-fluid">
      <div className="row flex-nowrap">
        <div className="col-auto col-md-3 col-xl-2 px-sm-2 px-0 bg-dark">
          <Sidebar/>
        </div>
      </div>
    </div>
  )
}
