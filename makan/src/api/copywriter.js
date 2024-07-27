import axios from 'axios';
import Ajv from 'ajv';

const COPY_WRITER_URL =
  'https://asia-east1-personal-299420.cloudfunctions.net/CopyWriteFood';

const schema = {
  type: 'object',
  properties: {
    name: { type: 'string' },
  },
  required: ['name'],
  additionalProperties: false,
};

const ajv = new Ajv();
const validate = ajv.compile(schema);

const fetchFoodCopyWrite = async (data) => {
  const valid = validate(data);

  if (!valid) {
    // Handle validation errors
    console.error('Validation errors:', validate.errors);
    throw new Error('Invalid data format');
  }

  try {
    const response = await axios.post(COPY_WRITER_URL, data, {});
    console.log('Success:', response.data);
    return response.data;
  } catch (error) {
    console.error('Error:', error);
    throw error;
  }
};

export default fetchFoodCopyWrite;
