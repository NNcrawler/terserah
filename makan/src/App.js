import './App.css';

import { useGeolocation } from '@uidotdev/usehooks';
import fetchFoodCopyWrite from './api/copywriter';
import { useEffect, useState } from 'react';
import fetchRecommendations from './api/recommendation';
import Card from './components/card';
import SlideShow from './components/slider';

function App() {
  const location = useGeolocation();

  const [recommendations, setRecommendations] = useState([]);
  useEffect(() => {
    const latitude = location.latitude;
    const longitude = location.longitude;
    if (latitude && longitude) {
      fetchRecommendations({
        longitude: longitude,
        latitude: latitude,
      }).then((response) => {
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
      <div className="flex items-center justify-center h-screen">
        <div className="max-w-sm h-screen">
          {copyWrite && (
            <SlideShow slides={recommendations}>
              {recommendations.map((recommendation, i) => (
                <Card
                  key={i}
                  title={recommendation.name}
                  content={<p>{copyWrite}</p>}
                  location={recommendation.location}
                />
              ))}
            </SlideShow>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
