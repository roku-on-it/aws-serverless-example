import { APIGatewayProxyEvent, APIGatewayProxyResult } from 'aws-lambda';
import { lambdaHandler } from '../app';
import { describe, expect, it } from '@jest/globals';

describe('Unit test for app handler', function () {
  it('verifies successful response', async () => {
    const result: APIGatewayProxyResult = await lambdaHandler({} as APIGatewayProxyEvent);

    expect(result.statusCode).toEqual(200);
  });
});
