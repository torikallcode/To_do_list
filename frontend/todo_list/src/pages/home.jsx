import React, { useState, useEffect } from 'react';
import axios from 'axios';

export const Home = () => {
  const [list, setList] = useState([]); // Pastikan ini adalah array
  const [input, setInput] = useState('');
  const [edit, setEdit] = useState(false);
  const [id, setId] = useState(null);

  const api = axios.create({
    baseURL: 'http://localhost:8080', // Perbaiki typo
  });

  const fetchList = async () => {
    try {
      const response = await api.get('/list');
      // Periksa apakah data adalah array
      if (Array.isArray(response.data)) {
        setList(response.data);
      } else {
        console.error('Data yang diterima bukan array:', response.data);
        setList([]); // Reset list jika data tidak valid
      }
    } catch (error) {
      console.log('Error fetching data', error);
    }
  };

  useEffect(() => {
    fetchList();
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await api.post('/list', { name_list: input }); // Pastikan format data sesuai
      setList([...list, response.data]);
      setInput('');
    } catch (error) {
      console.log('Error fetching data', error);
    }
  };

  const handleDelete = async (id) => {
    try {
      await api.delete(`/list/${id}`);
      setList(list.filter((item) => item.id !== id));
    } catch (error) {
      console.log('Error fetching data', error);
    }
  };

  const handleEdit = (list) => {
    setEdit(true);
    setId(list.id);
    setInput(list.name_list);
  };

  const handleUpdate = async (e) => {
    e.preventDefault();
    try {
      const response = await api.put(`/list/${id}`, { name_list: input }); // Pastikan format data sesuai
      setList(list.map(item => item.id === id ? response.data : item));
      setInput('');
      setEdit(false);
      setId(null)
    } catch (error) {
      console.log('Error fetching data', error);
    }
  };

  const handleToggleStatus = async (itemId) => {
    try {
      const response = await api.patch(`/list/${itemId}/status`);

      // Update list dengan status baru
      setList(list.map(item =>
        item.id === itemId
          ? { ...item, status: response.data.status }
          : item
      ));
    } catch (error) {
      console.error('Error toggling status', error);
    }
  };

  return (
    <div className='flex flex-col items-center justify-center pt-10'>
      <div className='bg-sky-600 max-w-lg w-full flex p-2 space-y-3 rounded-md flex-col'>
        <h1 className='text-2xl font-bold text-white text-center'>Todo-List</h1>
        <form
          onSubmit={edit ? handleUpdate : handleSubmit}
          className="flex border-2 border-sky-900 rounded-xl w-full"
          aria-label="simple-form"
        >
          <div className="flex-1">
            <input
              id='name_list'
              value={input}
              onChange={(e) => setInput(e.target.value)}
              type="text"
              placeholder="Enter your content"
              className="w-full p-3 bg-transparent outline-none font-semibold text-slate-100"
            />
          </div>
          <button
            className="flex-shrink-0 p-3 text-white bg-sky-900 rounded-lg"
            type='submit'
          >
            Submit
          </button>
        </form>
        <ul>
          {Array.isArray(list) && list.map((item) => (
            <li
              key={item.id}
              className={`
                flex w-full justify-between space-x-5 items-center 
                rounded-md py-2 px-3 
                ${item.status ? 'bg-green-600' : 'bg-sky-900'}
              `}
            >
              {/* Tambahkan checkbox */}
              <input
                type="checkbox"
                checked={item.status}
                onChange={() => handleToggleStatus(item.id)}
                className="form-checkbox h-5 w-5 text-green-500"
              />

              <span
                className={`
                  overflow-hidden whitespace-nowrap text-ellipsis max-w-xs 
                  ${item.status ? 'line-through text-gray-300' : 'text-white'}
                `}
              >
                {item.name_list}
              </span>

              <div className='flex space-x-2'>
                <button
                  className="inline-flex items-center justify-center px-3 py-2 font-sans text-white bg-sky-500 rounded-lg cursor-pointer"
                  onClick={() => handleEdit(item)}
                >
                  Edit
                </button>
                <button
                  className="inline-flex items-center justify-center px-3 py-2 font-sans text-white bg-red-500 rounded-lg cursor-pointer"
                  onClick={() => handleDelete(item.id)}
                >
                  Hapus
                </button>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};
