import axios from 'axios';
import Ajv from 'ajv';

const COPY_WRITER_URL =
  'https://asia-east1-personal-299420.cloudfunctions.net/CopyWriteFood';

const ajv = new Ajv();
const validatorMap = {
  localGuideRecommendation: ajv.compile({
    type: 'object',
    properties: {
      name: { type: 'string' },
    },
    required: ['name'],
    additionalProperties: false,
  }),
  reviewSummarizer: ajv.compile({
    type: 'object',
    properties: {
      reviews: {
        type: 'array',
        items: {
          type: 'string',
        },
      },
    },
    required: ['reviews'],
    additionalProperties: false,
  }),
};

const copyWriter = async (mode, data) => {
  if (validatorMap[mode] === undefined) {
    throw new Error('Invalid mode');
  }

  const validate = validatorMap[mode];

  const valid = validate(data);

  if (!valid) {
    // Handle validation errors
    throw new Error(
      `Invalid data format ${JSON.stringify(
        validate.errors
      )}\ngot ${JSON.stringify(data)}`
    );
  }

  try {
    const response = await axios.post(COPY_WRITER_URL, {
      mode,
      data,
    });
    console.log('Success:', response.data);
    return response.data;
  } catch (error) {
    console.error('Error:', error);
    throw error;
  }
};

export default copyWriter;
