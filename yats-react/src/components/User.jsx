
import React, { useEffect, useState } from 'react';
import ReactDom from 'react-dom/client';

const url = `http://localhost:4001/admin`//?${param}`

const User = () =>{

  const [data, setData] = useState(null)
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);


  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(url);
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const result = await response.json();
        setData(result);
      } catch (err) {
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

if (loading) {
    return <div>Loading data...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <div>
      {
        data.map((item, i) => {
          console.log(i,item);
          return <tr className={`item1 ${i}`}>
              <td className="t-op-nextlvl">{item.FirstName}                         </td>
              <td className="t-op-nextlvl">{item.LastName}                          </td>
              <td className="t-op-nextlvl">{item.Role}                              </td>
              <td className="t-op-nextlvl ">{(item.Status) ? "Active" : "Inactive"} </td>
              <td><button className="t-op-nextlvl ">Edit</button></td>
            </tr>
          
        })
      }
    </div>
  );
}

export default User;