import './App.css';
import LocationFetcher from './LocationFetcher';
import { useGeolocation } from '@uidotdev/usehooks';
import fetchFoodCopyWrite from './api/copywriter';
import { useEffect, useState } from 'react';
import fetchRecommendations from './api/recommendation';
import Card from './components/card';

function App() {
  const location = useGeolocation();

  const [recommendations, setRecommendations] = useState([]);
  useEffect(() => {
    if (location.latitude && location.longitude) {
      fetchRecommendations(location).then((response) => {
        setRecommendations(response.data);
      });
    }
  }, [location]);

  const [copyWrite, setCopyWrite] = useState('');
  useEffect(() => {
    if (recommendations.length > 0) {
      fetchFoodCopyWrite({ name: recommendations[0].name }).then(
        (copyWrite) => {
          setCopyWrite(copyWrite);
        }
      );
    }
  }, [recommendations]);

  return (
    <div className="App">
      <div class="flex items-center justify-center h-screen">
        <div>
          {/* <LocationFetcher location={location} /> */}
          {/* {recommendations.length > 0 &&
            recommendations.map((recommendation, i) => (
              <p id={i}>{recommendation.name}</p>
            ))} */}
          {copyWrite && <Card title={recommendations[0].name} content={<p>{copyWrite}</p>} />}
        </div>
      </div>
    </div>
  );
}

export default App;
