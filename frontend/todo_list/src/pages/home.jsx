import react from 'react'

export const Home = () => {
  return (
    <div className='flex flex-col items-center justify-center pt-10'>
      <div className='bg-sky-600 max-w-lg w-full flex p-2 space-y-3 rounded-md flex-col '>
        <h1 className='text-2xl font-bold text-white text-center'>Todo-List</h1>
        <form
          className="flex border-2 border-sky-900 rounded-xl w-full"
          aria-label="simple-form"
        >
          <div className="flex-1">
            <input
              type="text"
              placeholder="Enter your content"
              className="w-full p-3 bg-transparent outline-none font-semibold text-slate-100"
            />
          </div>
          <button className="flex-shrink-0 p-3 text-white bg-sky-900 rounded-lg">
            Submit
          </button>
        </form>
        <ul>
          <li className='flex w-full justify-between space-x-5 items-center bg-sky-900 rounded-md py-2 px-3'>
            <span className=' text-green-300'>selesai</span>
            <span className='overflow-hidden whitespace-nowrap text-ellipsis max-w-xs text-white'>Lorem ipsum, dolor sit amet consectetur adipisicing elit. raesentium.</span>
            <div className='flex space-x-2'>
              <button className="inline-flex items-center justify-center px-3 py-2 font-sans text-white bg-sky-500 rounded-lg cursor-pointer ">
                Edit
              </button>
              <button className="inline-flex items-center justify-center px-3 py-2 font-sans text-white bg-red-500 rounded-lg cursor-pointer ">
                Hapus
              </button>
            </div>
          </li>
        </ul>

      </div>
    </div>
  )
}