import axios from 'axios';
import Ajv from 'ajv';
// function fetchRecommendationsSimulation({ latitude = 0, longitude = 0 }) {
//   return new Promise((resolve, reject) => {
//     setTimeout(() => {
//       const response = {
//         data: [
//           {
//             name: 'Bakso Joni Blok S',
//             dishType: 'bakso',
//             location: {
//               googleMaps: 'https://goo.gl/maps/1J1J1J1J1J1J1J1J1',
//               address: 'Jl. Blok S No. 1, Jakarta Selatan',
//             },
//           },
//           {
//             name: 'Nasi Goreng Kambing',
//             dishType: 'nasi goreng',
//             location: {
//               googleMaps: 'https://goo.gl/maps/2J2J2J2J2J2J2J2J2',
//               address: 'Jl. Kambing No. 2, Jakarta Selatan',
//             },
//           },
//           {
//             name: 'Sate Ayam Madura',
//             dishType: 'sate',
//             location: {
//               googleMaps: 'https://goo.gl/maps/3J3J3J3J3J3J3J3J3',
//               address: 'Jl. Madura No. 3, Jakarta Selatan',
//             },
//           },
//         ],
//       };
//       resolve(response);
//     }, 30); // Simulating a delay of 1 second
//   });
// }

const RECOMMENDER_URL =
  'https://asia-southeast1-exploration-way.cloudfunctions.net/GetRecommendations';

const schema = {
  type: 'object',
  properties: {
    longitude: { type: 'number' },
    latitude: { type: 'number' },
  },
  required: ['longitude', 'latitude'],
  additionalProperties: false,
};

const ajv = new Ajv();
const validate = ajv.compile(schema);

async function fetchRecommendations(data) {
  const valid = validate(data);

  if (!valid) {
    // Handle validation errors
    console.error('Validation errors:', validate.errors);
    throw new Error('Invalid data format');
  }

  try {
    const response = await axios.get(
      `${RECOMMENDER_URL}?latitude=${data.latitude}&longitude=${data.longitude}`,
      data,
      {}
    );
    return response.data;
  } catch (error) {
    throw error;
  }
}

export default fetchRecommendations;
