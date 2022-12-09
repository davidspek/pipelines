# coding: utf-8

"""
    Kubeflow Pipelines API

    This file contains REST API specification for Kubeflow Pipelines. The file is autogenerated from the swagger definition.

    Contact: kubeflow-pipelines@google.com
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import kfp_server_api
from kfp_server_api.models.v2beta1_pipeline_task_detail import V2beta1PipelineTaskDetail  # noqa: E501
from kfp_server_api.rest import ApiException

class TestV2beta1PipelineTaskDetail(unittest.TestCase):
    """V2beta1PipelineTaskDetail unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V2beta1PipelineTaskDetail
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = kfp_server_api.models.v2beta1_pipeline_task_detail.V2beta1PipelineTaskDetail()  # noqa: E501
        if include_optional :
            return V2beta1PipelineTaskDetail(
                run_id = '0', 
                task_id = '0', 
                display_name = '0', 
                create_time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                start_time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                end_time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                executor_detail = kfp_server_api.models.v2beta1_pipeline_task_executor_detail.v2beta1PipelineTaskExecutorDetail(
                    main_job = '0', 
                    pre_caching_check_job = '0', 
                    failed_main_jobs = [
                        '0'
                        ], 
                    failed_pre_caching_check_jobs = [
                        '0'
                        ], ), 
                state = 'RUNTIMESTATE_UNSPECIFIED', 
                execution_id = '0', 
                error = kfp_server_api.models.rpc_status.rpcStatus(
                    code = 56, 
                    message = '0', 
                    details = [
                        kfp_server_api.models.protobuf_any.protobufAny(
                            type_url = '0', 
                            value = 'YQ==', )
                        ], ), 
                inputs = {
                    'key' : kfp_server_api.models.v2beta1_artifact_list.v2beta1ArtifactList(
                        artifact_ids = [
                            '0'
                            ], )
                    }, 
                outputs = {
                    'key' : kfp_server_api.models.v2beta1_artifact_list.v2beta1ArtifactList(
                        artifact_ids = [
                            '0'
                            ], )
                    }, 
                parent_task_id = '0', 
                state_history = [
                    kfp_server_api.models.v2beta1_runtime_status.v2beta1RuntimeStatus(
                        update_time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        state = 'RUNTIMESTATE_UNSPECIFIED', 
                        error = kfp_server_api.models.rpc_status.rpcStatus(
                            code = 56, 
                            message = '0', 
                            details = [
                                kfp_server_api.models.protobuf_any.protobufAny(
                                    type_url = '0', 
                                    value = 'YQ==', )
                                ], ), )
                    ]
            )
        else :
            return V2beta1PipelineTaskDetail(
        )

    def testV2beta1PipelineTaskDetail(self):
        """Test V2beta1PipelineTaskDetail"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
